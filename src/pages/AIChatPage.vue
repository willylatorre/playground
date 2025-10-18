<script setup lang="ts">
import { ref } from 'vue'
import { useChat } from '@/composables/useChat'

const { messages, status, sendMessage } = useChat()
const input = ref('')

const handleSubmit = async (event?: Event) => {
  event?.preventDefault()
  const prompt = input.value
  if (!prompt.trim()) return
  await sendMessage(prompt)
  input.value = ''
}
</script>

<template>
  <div class="space-y-4">
    <!-- Page Header -->
    <div>
      <h1 class="text-3xl font-bold text-slate-900 mb-2">AI Chat</h1>
      <p class="text-slate-600 max-w-2xl space-y-2">
        <span>
          Explore interactions with GenAI powered by a Golang server and OpenAI. This is a learning
          project to understand streaming responses, clean architecture, and AI integration. Try
          chatting and see if you can find any easter eggs! 🥚
        </span>
        <span>
          Using the
          <a
            class="font-medium text-slate-900 underline decoration-slate-300 underline-offset-4 hover:text-slate-700"
            href="https://github.com/openai/openai-go"
            target="_blank"
            rel="noopener"
          >
            official openai-go SDK
          </a>
          under the hood.
        </span>
      </p>
    </div>

    <!-- Chat Palette -->
    <div class="mt-[3rem]">
      <UChatPalette class="h-[28rem] border border-slate-200 rounded overflow-hidden">
        <UChatMessages :messages="messages" :status="status" :should-auto-scroll="true" />

        <template #prompt>
          <UContainer class="flex flex-col gap-3 py-4 bg-slate-100">
            <UChatPrompt
              v-model="input"
              placeholder="Ask me anything about Go, Vue, or try to find easter eggs..."
              @submit="handleSubmit"
            >
              <UChatPromptSubmit :status="status" />
            </UChatPrompt>
            <div class="flex items-center gap-2">
              <UIcon name="i-lucide-bot" class="w-5 h-5 text-slate-400" />
              <span class="text-sm text-slate-500">GPT-5</span>
            </div>
          </UContainer>
        </template>
      </UChatPalette>
    </div>

    <!-- Takeaways -->
    <section class="space-y-6 mt-[6rem] mb-[6rem]">
      <div>
        <h2 class="text-2xl font-semibold text-slate-900">Takeaways</h2>
        <p class="text-slate-600">
          Notes from building this chat experience with Go, Nuxt, and the OpenAI SDK.
        </p>
      </div>

      <div class="grid grid-cols-1 gap-4 md:grid-cols-3">
        <UCard>
          <template #header>
            <div class="flex items-center gap-2 text-slate-900">
              <UIcon name="i-lucide-puzzle" class="w-5 h-5" />
              <span class="font-medium">SDK Flexibility</span>
            </div>
          </template>
          <p class="text-slate-600">
            Integration with the Go SDK is straightforward but less flexible compared to the
            TypeScript AI SDK. The TS version still feels like the most adaptable option for rapid
            prototyping and feature parity.
          </p>
        </UCard>

        <UCard>
          <template #header>
            <div class="flex items-center gap-2 text-slate-900">
              <UIcon name="i-lucide-layers" class="w-5 h-5" />
              <span class="font-medium">Architecture Guidance</span>
            </div>
          </template>
          <p class="text-slate-600">
            The SDK repository lacks guidance on clean architecture patterns. Moving the client into
            a service layer with clear abstractions keeps the codebase easier to evolve and test.
          </p>
        </UCard>

        <UCard>
          <template #header>
            <div class="flex items-center gap-2 text-slate-900">
              <UIcon name="i-lucide-message-square-text" class="w-5 h-5" />
              <span class="font-medium">UI Message Shape</span>
            </div>
          </template>
          <p class="text-slate-600">
            UI libraries often expect message-specific structures. Adapting responses with a `parts`
            array (mirroring the AI SDK) makes it easier to plug into Nuxt UI components without
            bespoke renderers.
          </p>
        </UCard>
      </div>
    </section>
  </div>
</template>
