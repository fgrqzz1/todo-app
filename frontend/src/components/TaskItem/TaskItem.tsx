import type { Task } from '@/types/task'
import styles from './TaskItem.module.css'

interface TaskItemProps {
  task: Task
  onToggleDone: (id: number) => void
  onDelete: (id: number) => void
  isUpdating?: boolean
}

function formatDate(iso: string) {
  const d = new Date(iso)
  const now = new Date()
  const isToday = d.toDateString() === now.toDateString()
  if (isToday) {
    return d.toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
  }
  return d.toLocaleDateString('ru-RU', { day: 'numeric', month: 'short' })
}

export function TaskItem({ task, onToggleDone, onDelete, isUpdating }: TaskItemProps) {
  return (
    <article
      className={`${styles.card} ${task.Done ? styles.cardDone : ''}`}
      data-task-id={task.ID}
    >
      <label className={styles.checkboxWrap}>
        <input
          type="checkbox"
          checked={task.Done}
          onChange={() => onToggleDone(task.ID)}
          disabled={isUpdating}
          className={styles.checkbox}
        />
        <span className={styles.checkmark} />
      </label>
      <div className={styles.content}>
        <span className={styles.title}>{task.Title}</span>
        {task.Description ? (
          <p className={styles.description}>{task.Description}</p>
        ) : null}
        <time className={styles.date} dateTime={task.CreatedAt}>
          {formatDate(task.CreatedAt)}
        </time>
      </div>
      <button
        type="button"
        className={styles.delete}
        onClick={() => onDelete(task.ID)}
        disabled={isUpdating}
        title="Удалить"
        aria-label="Удалить задачу"
      >
        ×
      </button>
    </article>
  )
}
