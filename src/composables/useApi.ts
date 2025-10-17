import { ref, type Ref } from 'vue'

export interface Coffee {
  id: number
  counter: number
  last_update: string
}

const API_BASE_URL = 'http://localhost:8080/api'

export function useApi() {
  const loading: Ref<boolean> = ref(false)
  const error: Ref<string | null> = ref(null)

  const checkHealth = async (): Promise<{ status: string; message: string } | null> => {
    try {
      const response = await fetch(`${API_BASE_URL}/health`)
      if (!response.ok) {
        return null
      }
      return await response.json()
    } catch {
      return null
    }
  }

  const getCoffee = async (): Promise<Coffee | null> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/coffee`)
      const result = await response.json()
      return result.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch coffee counter'
      return null
    } finally {
      loading.value = false
    }
  }

  const incrementCoffee = async (): Promise<Coffee | null> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/coffee/increment`, {
        method: 'POST',
      })
      const result = await response.json()
      return result.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to increment coffee counter'
      return null
    } finally {
      loading.value = false
    }
  }

  return {
    loading,
    error,
    checkHealth,
    getCoffee,
    incrementCoffee,
  }
}
