package websocket

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Connection 表示一个WebSocket连接
type Connection struct {
	conn *websocket.Conn
	send chan []byte
}

// ConnectionPool 管理所有的WebSocket连接
type ConnectionPool struct {
	connections map[*Connection]bool
	broadcast   chan []byte
	register    chan *Connection
	unregister  chan *Connection
	mutex       sync.RWMutex
	done        chan struct{}
}

// NewConnectionPool 创建一个新的连接池
func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[*Connection]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		done:        make(chan struct{}),
	}
}

// Run 运行连接池
func (pool *ConnectionPool) Run() {
	for {
		select {
		case conn := <-pool.register:
			pool.mutex.Lock()
			pool.connections[conn] = true
			pool.mutex.Unlock()

		case conn := <-pool.unregister:
			if _, ok := pool.connections[conn]; ok {
				pool.mutex.Lock()
				delete(pool.connections, conn)
				close(conn.send)
				pool.mutex.Unlock()
			}

		case message := <-pool.broadcast:
			pool.mutex.RLock()
			for conn := range pool.connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(pool.connections, conn)
				}
			}
			pool.mutex.RUnlock()

		case <-pool.done:
			pool.mutex.Lock()
			for conn := range pool.connections {
				close(conn.send)
				conn.conn.Close()
			}
			pool.connections = make(map[*Connection]bool)
			pool.mutex.Unlock()
			return
		}
	}
}

// Broadcast 向所有连接广播消息
func (pool *ConnectionPool) Broadcast(message []byte) error {
	if len(pool.connections) == 0 {
		return errors.New("no connections in pool")
	}
	pool.broadcast <- message
	return nil
}

// HandleConnection 处理单个WebSocket连接
func (pool *ConnectionPool) HandleConnection(conn *websocket.Conn) {
	c := &Connection{
		conn: conn,
		send: make(chan []byte, 256),
	}

	pool.register <- c

	// 启动goroutine处理写操作
	go func() {
		defer func() {
			pool.unregister <- c
			c.conn.Close()
		}()

		for message := range c.send {
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("error writing message: %v", err)
				return
			}
		}
	}()

	// 处理读操作
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}
		// 收到消息后广播给所有连接
		pool.broadcast <- message
	}
}

// Shutdown 方法用于优雅关闭连接池
func (pool *ConnectionPool) Shutdown(ctx context.Context) {
	close(pool.done)
	// 等待所有连接关闭
	pool.mutex.RLock()
	defer pool.mutex.RUnlock()
	log.Printf("WebSocket connection pool shutdown, closed %d connections", len(pool.connections))
}
