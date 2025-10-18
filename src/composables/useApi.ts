import { ref, type Ref } from 'vue'
import type { Coffee } from '@/types/api-generated'

const API_BASE_URL = 'http://localhost:8080/api'

export function useApi() {
  const loading: Ref<boolean> = ref(false)
  const error: Ref<string | null> = ref(null)

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

  const sendChatMessage = async (
    message: string,
    onChunk: (chunk: string) => void,
    onComplete: () => void,
    onError: (error: string) => void
  ): Promise<void> => {
    try {
      const response = await fetch(`${API_BASE_URL}/chat/message`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message }),
      })

      if (!response.ok) {
        const errorData = await response.json()
        onError(errorData.error || 'Failed to send message')
        return
      }

      const reader = response.body?.getReader()
      const decoder = new TextDecoder()

      if (!reader) {
        onError('Failed to read response')
        return
      }

      while (true) {
        const { done, value } = await reader.read()
        if (done) break

        const chunk = decoder.decode(value)
        const lines = chunk.split('\n')

        for (const line of lines) {
          if (line.startsWith('data: ')) {
            const data = line.slice(6)
            if (data.trim()) {
              onChunk(data)
            }
          }
        }
      }

      onComplete()
    } catch (err) {
      onError(err instanceof Error ? err.message : 'Unknown error occurred')
    }
  }

  return {
    loading,
    error,
    getCoffee,
    incrementCoffee,
    sendChatMessage,
  }
}
