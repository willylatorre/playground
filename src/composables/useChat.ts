import { ref } from 'vue'
import type { ChatMessage } from '@/types/api-generated'
import { useApi } from './useApi'

type ChatRole = 'assistant' | 'user' | 'system'

type UIChatMessage = Omit<ChatMessage, 'role'> & {
  role: ChatRole
  parts: { type: 'text'; text: string }[]
}

const createTextPart = (text: string) => ({ type: 'text' as const, text })

const ensureFirstPart = (message: UIChatMessage) => {
  if (!message.parts.length) {
    message.parts.push(createTextPart(''))
  }
  return message.parts[0]
}

export function useChat() {
  const { sendChatMessage } = useApi()

  const messages = ref<UIChatMessage[]>([
    {
      id: '1',
      role: 'assistant',
      content:
        "Hey! 👋 I'm an AI assistant built with Go and OpenAI. I'm here to help you learn about web development, Go, Vue, and clean architecture. Feel free to ask me anything, or try to find some easter eggs! 🥚",
      parts: [
        createTextPart(
          "Hey! 👋 I'm an AI assistant built with Go and OpenAI. I'm here to help you learn about web development, Go, Vue, and clean architecture. Feel free to ask me anything, or try to find some easter eggs! 🥚",
        ),
      ],
      timestamp: new Date(Date.now() - 60000).toISOString(),
    },
    {
      id: '2',
      role: 'user',
      content: 'What is this project about?',
      parts: [createTextPart('What is this project about?')],
      timestamp: new Date(Date.now() - 45000).toISOString(),
    },
    {
      id: '3',
      role: 'assistant',
      content:
        'This is a learning playground exploring the intersection of Go backends and modern Vue frontends. It demonstrates clean architecture, type safety with TypeScript, SQLite databases, and AI integration with OpenAI streaming. The entire codebase is designed to be educational and showcase best practices!',
      parts: [
        createTextPart(
          'This is a learning playground exploring the intersection of Go backends and modern Vue frontends. It demonstrates clean architecture, type safety with TypeScript, SQLite databases, and AI integration with OpenAI streaming. The entire codebase is designed to be educational and showcase best practices!',
        ),
      ],
      timestamp: new Date(Date.now() - 30000).toISOString(),
    },
  ])

  const status = ref<'ready' | 'submitted' | 'streaming' | 'error'>('ready')
  const currentAssistantMessage = ref('')

  const sendMessage = async (prompt: string): Promise<void> => {
    if (!prompt.trim()) return

    // Clone current history before adding the new user message
    const history: ChatMessage[] = messages.value.map(({ id, role, content, timestamp }) => ({
      id,
      role,
      content,
      timestamp,
    }))

    // Add user message
    const userMessage: UIChatMessage = {
      id: Date.now().toString(),
      role: 'user',
      content: prompt,
      parts: [createTextPart(prompt)],
      timestamp: new Date().toISOString(),
    }
    messages.value.push(userMessage)

    // Prepare assistant message
    const assistantMessageId = (Date.now() + 1).toString()
    const assistantMessage: UIChatMessage = {
      id: assistantMessageId,
      role: 'assistant',
      content: '',
      parts: [createTextPart('')],
      timestamp: new Date().toISOString(),
    }

    const assistantTextPart = ensureFirstPart(assistantMessage)

    messages.value.push(assistantMessage)
    currentAssistantMessage.value = ''
    status.value = 'submitted'

    // Send message and stream response
    await sendChatMessage(
      history,
      prompt,
      (chunk: string) => {
        // Append chunk to current message
        status.value = 'streaming'
        currentAssistantMessage.value += chunk
        assistantMessage.content = currentAssistantMessage.value
        if (assistantTextPart) {
          assistantTextPart.text = currentAssistantMessage.value
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
        assistantMessage.content = `Error: ${error}`
        if (assistantTextPart) {
          assistantTextPart.text = `Error: ${error}`
        }
      },
    )
  }

  const clearMessages = () => {
    messages.value = []
  }

  return {
    messages,
    status,
    sendMessage,
    clearMessages,
  }
}
