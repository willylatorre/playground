<script setup lang="ts">
// AI Chat page - blank skeleton for now
</script>

<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <UPageHero
      title="AI Chat"
      description="Have a conversation with AI - coming soon!"
    />

    <!-- Chat Interface Skeleton -->
    <UCard>
      <template #header>
        <h3 class="text-xl font-semibold">Chat Interface</h3>
      </template>

      <div class="space-y-4">
        <!-- Chat Messages Area -->
        <div class="min-h-96 bg-slate-50 rounded-lg p-4 border-2 border-dashed border-slate-200 flex items-center justify-center">
          <div class="text-center text-slate-500">
            <UIcon name="i-heroicons-chat-bubble-left-right" class="w-12 h-12 mx-auto mb-3" />
            <h4 class="font-medium mb-2">Chat Interface Coming Soon</h4>
            <p class="text-sm">This will be an interactive AI chat experience</p>
          </div>
        </div>

        <!-- Message Input -->
        <div class="flex gap-3">
          <UInput
            placeholder="Type your message here..."
            disabled
            class="flex-1"
          />
          <UButton disabled>
            <UIcon name="i-heroicons-paper-airplane" class="w-4 h-4" />
          </UButton>
        </div>

        <!-- Feature Preview -->
        <UAlert color="info" variant="subtle">
          <UIcon name="i-heroicons-information-circle" class="w-5 h-5 mr-3" />
          <div>
            <p class="font-medium">Upcoming Features:</p>
            <ul class="text-sm mt-1 list-disc list-inside">
              <li>Real-time AI conversations</li>
              <li>Context-aware responses</li>
              <li>Chat history and memory</li>
              <li>Multiple AI personalities</li>
            </ul>
          </div>
        </UAlert>
      </div>
    </UCard>
  </div>
</template>
