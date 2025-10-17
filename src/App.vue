<script setup lang="ts">
import { ref, computed } from 'vue'

const searchTerm = ref('')
const adrianStatus = ref('Available')
const statusColors = ['primary', 'secondary', 'success', 'info', 'warning', 'error'] as const
const randomStatus = computed(() => statusColors[Math.floor(Math.random() * statusColors.length)])

// Search groups for CommandPalette
const searchGroups = [
  {
    id: 'about-me',
    label: 'About Me',
    items: [
      {
        label: 'About Me',
        suffix: 'My profile',
        to: '/',
        icon: 'i-lucide-user',
      },
      {
        label: 'Contact',
        suffix: 'adrian@example.com',
        action: () => {
          window.open('mailto:adrian@example.com', '_blank')
        },
        icon: 'i-lucide-mail',
      },
      {
        label: 'GitHub',
        suffix: '@adrianlatorre',
        action: () => {
          window.open('https://github.com/adrianlatorre', '_blank')
        },
        icon: 'i-lucide-github',
      },
    ],
  },
  {
    id: 'playground',
    label: 'Playground',
    items: [
      {
        label: 'AI Chat',
        suffix: 'Interactive conversation',
        to: '/ai-chat',
        icon: 'i-lucide-message-circle',
      },
    ],
  },
]

const links = [
  {
    label: 'About Me',
    icon: 'i-lucide-user',
    to: '/',
  },
  {
    label: 'Playground',
    icon: 'i-lucide-command',
    defaultOpen: true,
    children: [
      {
        label: 'AI Chat',
        icon: 'i-lucide-message-circle',
        to: '/ai-chat',
      },
    ],
  },
  {
    label: 'Settings',
    icon: 'i-lucide-settings',
    to: '/settings',
  },
]

const changeStatus = () => {
  const statuses = [
    'Available',
    'Deep in Code',
    'Coffee Break',
    'Debugging',
    'Having Ideas',
    'Procrastinating',
    'Being Awesome',
    'Thinking Deeply',
  ] as const
  adrianStatus.value = statuses[Math.floor(Math.random() * statuses.length)] || 'Available'
}
</script>

<template>
  <UApp>
    <UDashboardGroup>
      <!-- Sidebar -->
      <UDashboardSidebar>
        <!-- Header -->
        <template #header>
          <div class="flex items-center gap-3 px-4 py-3">
            <div class="w-8 h-8 bg-primary-500 rounded-lg flex items-center justify-center">
              <UIcon name="i-heroicons-code-bracket" class="w-5 h-5 text-white" />
            </div>
            <div class="flex-1 min-w-0">
              <h2 class="font-semibold text-slate-900 truncate">Adrian Latorre</h2>
              <p class="text-xs text-slate-500 truncate">v{{ Math.random().toFixed(2) }}</p>
            </div>
          </div>
        </template>

        <!-- Navigation -->
        <UDashboardSearch :groups="searchGroups" />
        <UNavigationMenu :items="links" orientation="vertical" />

        <!-- Footer -->
        <template #footer>
          <div class="px-4 py-3 border-t border-slate-200">
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 bg-slate-100 rounded-full flex items-center justify-center">
                <UIcon name="i-heroicons-user" class="w-4 h-4 text-slate-600" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-xs font-medium text-slate-900">Developer</p>
                <p class="text-xs text-slate-500 truncate">developer@example.com</p>
              </div>
            </div>
          </div>
        </template>
      </UDashboardSidebar>

      <!-- Main Panel -->
      <UDashboardPanel>
        <!-- Navbar -->
        <template #header>
          <UDashboardNavbar title="Adrian Latorre" :toggle="true">
            <template #right>
              <div class="flex items-center gap-4">
                <!-- Adrian Status with Popover -->
                <UPopover mode="click">
                  <UButton
                    @click="changeStatus"
                    :color="randomStatus"
                    variant="subtle"
                    size="sm"
                    square
                  >
                    <UIcon name="i-lucide-smile" class="w-4 h-4" />
                  </UButton>

                  <template #panel>
                    <div class="p-4 min-w-64">
                      <h4 class="font-semibold mb-3">System Status</h4>

                      <!-- Status Items -->
                      <div class="space-y-2">
                        <div class="flex items-center justify-between">
                          <span class="text-sm">API Services</span>
                          <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-green-500"></div>
                            <span class="text-xs text-green-600">Operational</span>
                          </div>
                        </div>

                        <div class="flex items-center justify-between">
                          <span class="text-sm">Database</span>
                          <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-green-500"></div>
                            <span class="text-xs text-green-600">Healthy</span>
                          </div>
                        </div>

                        <div class="flex items-center justify-between">
                          <span class="text-sm">File System</span>
                          <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-red-500"></div>
                            <span class="text-xs text-red-600">Issues</span>
                          </div>
                        </div>

                        <div class="flex items-center justify-between">
                          <span class="text-sm">Cache</span>
                          <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-green-500"></div>
                            <span class="text-xs text-green-600">Good</span>
                          </div>
                        </div>

                        <div class="flex items-center justify-between">
                          <span class="text-sm">External APIs</span>
                          <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-green-500"></div>
                            <span class="text-xs text-green-600">Connected</span>
                          </div>
                        </div>
                      </div>

                      <div class="border-t pt-3 mt-3">
                        <p class="text-xs text-slate-500">
                          Adrian's Status:
                          <span class="font-medium text-slate-700">{{ adrianStatus }}</span>
                        </p>
                      </div>
                    </div>
                  </template>
                </UPopover>

                <!-- Coffee Counter -->
                <div class="text-xs text-slate-500 flex items-center gap-1">
                  <UIcon name="i-lucide-coffee" class="w-3 h-3" />
                  {{ Math.floor(Math.random() * 15) + 5 }}
                </div>

                <!-- Random Fun Fact Button -->
                <UButton
                  icon="i-lucide-lightbulb"
                  size="sm"
                  square
                  variant="ghost"
                  color="neutral"
                />

                <!-- User Menu -->
                <UButton icon="i-lucide-user" size="sm" square variant="ghost" color="neutral" />
              </div>
            </template>
          </UDashboardNavbar>
        </template>

        <!-- Main Content -->
        <template #body>
          <UContainer>
            <RouterView />
          </UContainer>
        </template>
      </UDashboardPanel>
    </UDashboardGroup>
  </UApp>
</template>

<style scoped>
/* Dashboard specific styles can go here if needed */
</style>
