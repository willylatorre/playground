<script setup lang="ts">
import { ref } from 'vue'
import { useApi } from '@/composables/useApi'
import type { ChatMessage } from '@/types/api-generated'

const { sendChatMessage } = useApi()

const messages = ref<ChatMessage[]>([])
const status = ref<'ready' | 'submitted' | 'streaming' | 'error'>('ready')
const currentAssistantMessage = ref('')

const handleSubmit = async (prompt: string) => {
  if (!prompt.trim()) return

  // Add user message
  const userMessage: ChatMessage = {
    id: Date.now().toString(),
    role: 'user',
    content: prompt,
    timestamp: new Date().toISOString(),
  }
  messages.value.push(userMessage)

  // Prepare assistant message
  const assistantMessageId = (Date.now() + 1).toString()
  const assistantMessage: ChatMessage = {
    id: assistantMessageId,
    role: 'assistant',
    content: '',
    timestamp: new Date().toISOString(),
  }
  messages.value.push(assistantMessage)
  currentAssistantMessage.value = ''
  status.value = 'streaming'

  // Send message and stream response
  await sendChatMessage(
    prompt,
    (chunk: string) => {
      // Append chunk to current message
      currentAssistantMessage.value += chunk
      const msgIndex = messages.value.findIndex((m) => m.id === assistantMessageId)
      if (msgIndex !== -1 && messages.value[msgIndex]) {
        messages.value[msgIndex]!.content = currentAssistantMessage.value
      }
    },
    () => {
      // Streaming complete
      status.value = 'ready'
      currentAssistantMessage.value = ''
    },
    (error: string) => {
      // Error occurred
      status.value = 'error'
      const msgIndex = messages.value.findIndex((m) => m.id === assistantMessageId)
      if (msgIndex !== -1 && messages.value[msgIndex]) {
        messages.value[msgIndex]!.content = `Error: ${error}`
      }
    }
  )
}
</script>

<template>
  <div class="space-y-4">
    <!-- Page Header -->
    <div>
      <h1 class="text-3xl font-bold text-slate-900 mb-2">AI Chat</h1>
      <p class="text-slate-600 max-w-2xl">
        Explore interactions with GenAI powered by a Golang server and OpenAI. This is a learning
        project to understand streaming responses, clean architecture, and AI integration. Try
        chatting and see if you can find any easter eggs! 🥚
      </p>
    </div>

    <!-- Chat Palette -->
    <UChatPalette>
      <UChatMessages
        :messages="messages"
        :status="status"
        :should-auto-scroll="true"
      />

      <template #prompt>
        <UChatPrompt
          placeholder="Ask me anything about Go, Vue, or try to find easter eggs..."
          @submit="handleSubmit"
        />
      </template>
    </UChatPalette>
  </div>
</template>
