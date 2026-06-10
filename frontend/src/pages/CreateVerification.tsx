import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { api } from '../api/verification'
import { Notification } from '../components/Notification'
import { s } from '../styles'

export function CreateVerification() {
  const nav = useNavigate()
  const [email, setEmail] = useState('')
  const [loading, setLoading] = useState(false)
  const [success, setSuccess] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)

  const handle = async () => {
    if (!email) return
    setLoading(true)
    setSuccess(null)
    setError(null)
    try {
      await api.create(email, email)
      setSuccess('Письмо с подтверждением отправлено. Проверьте почту и перейдите по ссылке в письме.')
      setEmail('')
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={s.page}>
      <h2 style={s.title}>Верифицировать почту</h2>
      <div style={s.form}>
        <input
          style={s.input}
          type="email"
          placeholder="Введите email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          onKeyDown={e => e.key === 'Enter' && handle()}
        />
        <button style={s.btnPrimary} onClick={handle} disabled={loading || !email}>
          {loading ? 'Отправляю...' : 'Отправить письмо'}
        </button>
      </div>

      {success && <Notification type="success" message={success} />}
      {error   && <Notification type="error"   message={error}   />}

      <button style={s.btnBack} onClick={() => nav('/')}>← Назад</button>
    </div>
  )
}
