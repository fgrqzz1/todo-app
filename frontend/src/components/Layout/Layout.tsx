import styles from './Layout.module.css'

interface LayoutProps {
  children: React.ReactNode
}

/** Обёртка страницы: фон, контейнер, единый стиль. */
export function Layout({ children }: LayoutProps) {
  return (
    <div className={styles.wrapper}>
      <div className={styles.container}>{children}</div>
    </div>
  )
}
