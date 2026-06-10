const BASE = '/api/v1/verification'

export interface Verification {
  id: string
  ext_id: string
  target: string
  status: number
  status_text: string
  expires_at?: string
  verified_at?: string
  created_at: string
}

export interface StatusResponse {
  found: boolean
  status?: number
  status_text?: string
  message: string
  verification?: Verification
}

export interface ApiError {
  error: string
}

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  })
  const data = await res.json()
  if (!res.ok) {
    throw new Error((data as ApiError).error || 'Неизвестная ошибка')
  }
  return data as T
}

export const api = {
  checkStatus: (email: string): Promise<StatusResponse> =>
    request(`${BASE}/status?email=${encodeURIComponent(email)}`),

  create: (extUserId: string, email: string): Promise<{ id: string }> =>
    request(`${BASE}/create`, {
      method: 'POST',
      body: JSON.stringify({ ext_user_id: extUserId, email }),
    }),

  retry: (extUserId: string, email: string): Promise<{ message: string }> =>
    request(`${BASE}/retry`, {
      method: 'POST',
      body: JSON.stringify({ ext_user_id: extUserId, email }),
    }),
}