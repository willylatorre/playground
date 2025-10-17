import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'About Me',
      component: () => import('../pages/DashboardPage.vue'),
    },
    {
      path: '/ai-chat',
      name: 'AI Chat',
      component: () => import('../pages/AIChatPage.vue'),
    },
    {
      path: '/items',
      name: 'Items',
      component: () => import('../pages/ItemsPage.vue'),
    },
    {
      path: '/settings',
      name: 'Settings',
      component: () => import('../pages/SettingsPage.vue'),
    },
  ],
})

export default router
