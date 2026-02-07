import type { Task } from '@/types/task'
import { TaskItem } from '@/components/TaskItem/TaskItem'
import styles from './TaskList.module.css'

interface TaskListProps {
  tasks: Task[]
  onToggleDone: (id: number) => void
  onDelete: (id: number) => void
  updatingId: number | null
}

export function TaskList({ tasks, onToggleDone, onDelete, updatingId }: TaskListProps) {
  if (tasks.length === 0) {
    return (
      <div className={styles.empty}>
        <div className={styles.emptyIcon}>◇</div>
        <p className={styles.emptyText}>Пока нет задач</p>
        <p className={styles.emptyHint}>Добавьте задачу в форме справа</p>
      </div>
    )
  }

  return (
    <ul className={styles.list} role="list">
      {tasks.map((task) => (
        <li key={task.ID}>
          <TaskItem
            task={task}
            onToggleDone={onToggleDone}
            onDelete={onDelete}
            isUpdating={updatingId === task.ID}
          />
        </li>
      ))}
    </ul>
  )
}
