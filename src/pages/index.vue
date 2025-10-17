<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Define proper types
interface Item {
  id: string
  name: string
  data: string
  created_at: string
  updated_at: string
}

interface FormData {
  name: string
  data: string
}

const items = ref<Item[]>([])
const serverStatus = ref<boolean>(false)

// Form state
const showCreateForm = ref(false)
const editingItem = ref<Item | null>(null)
const formData = ref<FormData>({
  name: '',
  data: '',
})
const loading = ref(false)
const error = ref<string | null>(null)

// Load data on mount
onMounted(async () => {
  // TODO: Replace with useApi later
  serverStatus.value = false
  items.value = []
})

// Create new item
const handleCreate = async () => {
  if (!formData.value.name.trim() || !formData.value.data.trim()) return

  loading.value = true
  error.value = null

  try {
    // TODO: Replace with useApi later
    console.log('Creating item:', formData.value)
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 500))

    const newItem: Item = {
      id: Date.now().toString(),
      name: formData.value.name,
      data: formData.value.data,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    }

    items.value.push(newItem)
    formData.value = { name: '', data: '' }
    showCreateForm.value = false
  } catch {
    error.value = 'Failed to create item'
  } finally {
    loading.value = false
  }
}

// Start editing item
const startEditing = (item: Item) => {
  editingItem.value = item
  formData.value = {
    name: item.name,
    data: item.data,
  }
}

// Update item
const handleUpdate = async () => {
  if (!editingItem.value || !formData.value.name.trim() || !formData.value.data.trim()) return

  loading.value = true
  error.value = null

  try {
    // TODO: Replace with useApi later
    console.log('Updating item:', editingItem.value.id, formData.value)
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 500))

    const index = items.value.findIndex(item => item.id === editingItem.value!.id)
    if (index !== -1) {
      const existingItem = items.value[index]
      if (existingItem) {
        items.value[index] = {
          id: existingItem.id,
          name: formData.value.name,
          data: formData.value.data,
          created_at: existingItem.created_at,
          updated_at: new Date().toISOString(),
        }
      }
    }
    cancelEdit()
  } catch {
    error.value = 'Failed to update item'
  } finally {
    loading.value = false
  }
}

// Delete item
const handleDelete = async (id: string) => {
  if (confirm('Are you sure you want to delete this item?')) {
    loading.value = true
    error.value = null

    try {
      // TODO: Replace with useApi later
      console.log('Deleting item:', id)
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 500))

      items.value = items.value.filter((item) => item.id !== id)
    } catch {
      error.value = 'Failed to delete item'
    } finally {
      loading.value = false
    }
  }
}

// Cancel editing
const cancelEdit = () => {
  editingItem.value = null
  formData.value = { name: '', data: '' }
}

// Refresh items
const handleRefresh = async () => {
  loading.value = true
  error.value = null

  try {
    // TODO: Replace with useApi later
    console.log('Refreshing items')
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 500))
    // For now, just keep existing items
  } catch {
    error.value = 'Failed to refresh items'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
    <div class="container mx-auto px-4 py-8 max-w-4xl">
      <!-- Header -->
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-slate-800 mb-4">🚀 Hello Golang + Vue Playground</h1>
        <UBadge :color="serverStatus ? 'success' : 'error'" variant="subtle" size="lg">
          <UIcon name="i-heroicons-wifi" class="w-4 h-4 mr-2" />
          Backend: {{ serverStatus ? 'Online' : 'Offline' }}
        </UBadge>
      </div>

      <div class="space-y-8">
        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <UButton
            @click="showCreateForm = true"
            :disabled="loading"
            icon="i-heroicons-plus"
            size="lg"
          >
            Add New Item
          </UButton>
          <UButton
            @click="handleRefresh"
            :disabled="loading"
            icon="i-heroicons-arrow-path"
            size="lg"
            color="neutral"
          >
            Refresh
          </UButton>
        </div>

        <!-- Error Message -->
        <UAlert v-if="error" color="error" variant="subtle">
          <UIcon name="i-heroicons-exclamation-triangle" class="w-5 h-5 mr-3" />
          {{ error }}
        </UAlert>

        <!-- Create Form -->
        <UCard v-if="showCreateForm">
          <template #header>
            <h3 class="text-xl font-semibold">Create New Item</h3>
          </template>

          <form @submit.prevent="handleCreate" class="space-y-4">
            <UFormGroup label="Name" required>
              <UInput id="name" v-model="formData.name" placeholder="Enter item name" required />
            </UFormGroup>

            <UFormGroup label="Data" required>
              <UTextarea
                id="data"
                v-model="formData.data"
                placeholder="Enter item data"
                :rows="4"
                required
              />
            </UFormGroup>

            <div class="flex gap-3 pt-4">
              <UButton type="submit" :disabled="loading" size="lg">
                {{ loading ? 'Creating...' : 'Create Item' }}
              </UButton>
              <UButton
                type="button"
                @click="showCreateForm = false"
                color="neutral"
                size="lg"
              >
                Cancel
              </UButton>
            </div>
          </form>
        </UCard>

        <!-- Edit Form -->
        <UCard v-if="editingItem">
          <template #header>
            <h3 class="text-xl font-semibold">Edit Item</h3>
          </template>

          <form @submit.prevent="handleUpdate" class="space-y-4">
            <UFormGroup label="Name" required>
              <UInput id="edit-name" v-model="formData.name" required />
            </UFormGroup>

            <UFormGroup label="Data" required>
              <UTextarea id="edit-data" v-model="formData.data" :rows="4" required />
            </UFormGroup>

            <div class="flex gap-3 pt-4">
              <UButton type="submit" :disabled="loading" size="lg">
                {{ loading ? 'Updating...' : 'Update Item' }}
              </UButton>
              <UButton type="button" @click="cancelEdit" color="neutral" size="lg">
                Cancel
              </UButton>
            </div>
          </form>
        </UCard>

        <!-- Items Grid -->
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <h3 class="text-xl font-semibold text-slate-800">Items ({{ items.length }})</h3>
          </div>

          <!-- Loading State -->
          <div v-if="loading && items.length === 0" class="flex items-center justify-center py-12">
            <div class="flex items-center gap-3 text-slate-500">
              <UIcon name="i-heroicons-arrow-path" class="w-6 h-6 animate-spin" />
              Loading items...
            </div>
          </div>

          <!-- Empty State -->
          <div v-else-if="items.length === 0" class="text-center py-12">
            <div class="text-6xl mb-4">🎯</div>
            <h4 class="text-xl font-semibold text-slate-700 mb-2">No items yet</h4>
            <p class="text-slate-500">Create your first item above to get started!</p>
          </div>

          <!-- Items List -->
          <div v-else class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            <UCard
              v-for="item in items"
              :key="item.id"
              class="hover:shadow-lg transition-shadow duration-200"
            >
              <div class="flex items-start justify-between mb-4">
                <h4 class="text-lg font-semibold flex-1 pr-4">{{ item.name }}</h4>
                <div class="flex gap-1">
                  <UButton
                    @click="startEditing(item)"
                    icon="i-heroicons-pencil"
                    size="sm"
                    square
                    variant="ghost"
                    color="primary"
                  />
                  <UButton
                    @click="handleDelete(item.id)"
                    icon="i-heroicons-trash"
                    size="sm"
                    square
                    variant="ghost"
                    color="error"
                  />
                </div>
              </div>

              <p class="text-slate-600 mb-4 leading-relaxed">{{ item.data }}</p>

              <template #footer>
                <p class="text-xs text-slate-500">
                  <UIcon name="i-heroicons-clock" class="w-4 h-4 inline mr-1" />
                  Created: {{ new Date(item.created_at).toLocaleString() }}
                </p>
              </template>
            </UCard>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
