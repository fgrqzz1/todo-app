/** Модель задачи с бэкенда */
export interface Task {
  ID: number
  Title: string
  Description: string
  Done: boolean
  CreatedAt: string
}

/** Тело запроса на создание задачи */
export interface CreateTaskInput {
  Title: string
  Description?: string
  Done?: boolean
}
