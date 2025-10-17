import { ref, type Ref } from 'vue'

export interface Item {
  id: string
  name: string
  data: string
  created_at: string
  updated_at: string
}

export interface CreateItemInput {
  name: string
  data: string
}

const API_BASE_URL = 'http://localhost:8080/api'

export function useApi() {
  const loading: Ref<boolean> = ref(false)
  const error: Ref<string | null> = ref(null)

  const fetchItems = async (): Promise<Item[]> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/items`)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const result = await response.json()
      return result.data || []
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch items'
      return []
    } finally {
      loading.value = false
    }
  }

  const createItem = async (input: CreateItemInput): Promise<Item | null> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/items`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(input),
      })
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const result = await response.json()
      return result.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create item'
      return null
    } finally {
      loading.value = false
    }
  }

  const updateItem = async (id: string, input: Partial<CreateItemInput>): Promise<Item | null> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/items/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(input),
      })
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const result = await response.json()
      return result.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update item'
      return null
    } finally {
      loading.value = false
    }
  }

  const deleteItem = async (id: string): Promise<boolean> => {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_BASE_URL}/items/${id}`, {
        method: 'DELETE',
      })
      return response.ok
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete item'
      return false
    } finally {
      loading.value = false
    }
  }

  const checkHealth = async (): Promise<boolean> => {
    try {
      const response = await fetch(`${API_BASE_URL}/health`)
      return response.ok
    } catch {
      return false
    }
  }

  return {
    loading,
    error,
    fetchItems,
    createItem,
    updateItem,
    deleteItem,
    checkHealth,
  }
}
