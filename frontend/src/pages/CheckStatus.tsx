import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import type { StatusResponse } from '../api/verification'
import { api } from '../api/verification'
import { Notification } from '../components/Notification'
import { s } from '../styles'

export function CheckStatus() {
  const nav = useNavigate()
  const [email, setEmail] = useState('')
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState<StatusResponse | null>(null)
  const [error, setError] = useState<string | null>(null)

  const handle = async () => {
    if (!email) return
    setLoading(true)
    setError(null)
    setResult(null)
    try {
      const resp = await api.checkStatus(email)
      setResult(resp)
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  const notifType = () => {
    if (!result) return 'info'
    if (!result.found) return 'info'
    switch (result.status) {
      case 2: return 'success'   // verified
      case 3: return 'error'     // failed
      case 4: return 'warning'   // expired
      default: return 'info'     // pending
    }
  }

  return (
    <div style={s.page}>
      <h2 style={s.title}>Проверить статус верификации</h2>
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
          {loading ? 'Проверяю...' : 'Проверить'}
        </button>
      </div>

      {error && <Notification type="error" message={error} />}
      {result && <Notification type={notifType()} message={result.message} />}

      <button style={s.btnBack} onClick={() => nav('/')}>← Назад</button>
    </div>
  )
}
