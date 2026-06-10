interface Props {
  type: 'success' | 'error' | 'info' | 'warning'
  message: string
}

const colors: Record<Props['type'], string> = {
  success: '#d1fae5',
  error:   '#fee2e2',
  info:    '#dbeafe',
  warning: '#fef3c7',
}

const borders: Record<Props['type'], string> = {
  success: '#34d399',
  error:   '#f87171',
  info:    '#60a5fa',
  warning: '#fbbf24',
}

export function Notification({ type, message }: Props) {
  return (
    <div style={{
      marginTop: '1rem',
      padding: '1rem 1.25rem',
      borderRadius: '0.5rem',
      borderLeft: `4px solid ${borders[type]}`,
      background: colors[type],
      color: '#1f2937',
      fontSize: '0.95rem',
    }}>
      {message}
    </div>
  )
}