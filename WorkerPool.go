package pool

import (
  "errors"
  "sync"
)

type PoolTask interface {
  // Execute запускает выполнение задачи и возвращает nil,
  // либо возникшую ошибку.
  Execute() error
  // OnFailure будет обрабатывать ошибки, возникшие в  Execute(), то есть
  // пул должен вызывать OnFailure в случае, если Execute возвращает ошибку.
  OnFailure(error)
}

type WorkerPool interface {
  // Start подготавливает пул для обработки задач. Должен вызываться один раз
  // перед использованием пула. Очередные вызовы должны игнорироваться.
  Start()
  // Stop останавливает обработку в пуле. Должен вызываться один раз.
  // Очередные вызовы должны игнорироваться.
  Stop()
  // AddWork добавляет задачу для обработки пулом. Добавлять задачи
  // можно после вызова Start() и до вызова Stop().
  // Если на момент добавления в пуле нет
  // свободных ресурсов (очередь заполнена) -
  // эту функция ожидает их освобождения (либо вызова Stop).
  AddWork(PoolTask)
}

type MyPool struct {
  tasks      chan PoolTask
  numWorkers int
  onceStart  sync.Once
  onceStop   sync.Once
  stopFlag   chan struct{}
}

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error) {
  if numWorkers < 1 {
    return nil, errors.New("numWorkers не может быть меньше 1")
  }
  if channelSize < 0 {
    return nil, errors.New("channelSize не может быть меньше 0")
  }
  return &MyPool{
    tasks:      make(chan PoolTask, channelSize),
    numWorkers: numWorkers,
    onceStart:  sync.Once{},
    onceStop:   sync.Once{},
    stopFlag:   make(chan struct{}),
  }, nil
}

func (p *MyPool) Start() {
  p.onceStart.Do(func() {
    for i := 0; i < p.numWorkers; i++ {
      go func() {
        for {
          select {
          case <-p.stopFlag:
            return
          case task, ok := <-p.tasks:
            if !ok {
              return
            }
            if err := task.Execute(); err != nil {
              task.OnFailure(err)
            }
          }
        }
      }()
    }
  })
}

func (p *MyPool) Stop() {
  p.onceStop.Do(func() {
    close(p.stopFlag)
    close(p.tasks)
  })
}

func (p *MyPool) AddWork(task PoolTask) {
  for {
    select {
    case p.tasks <- task:
    case <-p.stopFlag:
    }
  }
}