<script setup lang="ts">
import AdrianStatus from './components/AdrianStatus.vue'
import CoffeeCounter from './components/CoffeeCounter.vue'
import FunFacts from './components/FunFacts.vue'

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
  {
    id: 'media',
    label: 'Pet projects',
    items: [
      {
        label: 'Projects & Videos',
        suffix: 'Personal work and presentations',
        to: '/media',
        icon: 'i-lucide-play-circle',
      },
    ],
  },
]

const links = [
  {
    label: 'Links',
    type: 'label',
  },
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
    label: 'Pet projects',
    icon: 'i-lucide-play-circle',
    to: '/media',
  },
  {
    label: 'Settings',
    icon: 'i-lucide-settings',
    to: '/settings',
  },
]
</script>

<template>
  <UApp>
    <UDashboardGroup>
      <!-- Sidebar -->
      <UDashboardSidebar :default-size="20">
        <!-- Header -->
        <template #header>
          <div class="flex items-center gap-3 mt-5">
            <div class="w-6 h-6 bg-primary-500 rounded-lg flex items-center justify-center">
              <UIcon name="i-heroicons-code-bracket" class="w-3 h-3 text-white" />
            </div>
            <div class="flex-1">
              <h2 class="font-semibold text-slate-900">Adrian Latorre</h2>
              <p class="text-xs text-slate-500 truncate">v1.37</p>
            </div>
          </div>
        </template>

        <!-- Navigation -->
        <UDashboardSearchButton class="mt-4" />
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
          <UDashboardNavbar title="Playground" :toggle="true">
            <template #right>
              <div class="flex items-center gap-4">
                <!-- Adrian Status -->
                <AdrianStatus />

                <!-- Coffee Counter -->
                <CoffeeCounter />

                <!-- Random Fun Fact Button -->
                <FunFacts />
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
