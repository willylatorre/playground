<script setup lang="ts">
import { ref } from 'vue'

const coffeeCount = ref(Math.floor(Math.random() * 15) + 5)

const toast = useToast()

// Crazy coffee messages that escalate in incredulity (every 10 clicks)
const coffeeMessages = [
  "That hit the spot! ☕",
  "You're really embracing the coffee life!",
  "Your energy levels are through the roof!",
  "Caffeine coursing through your veins!",
  "You're basically a coffee wizard now! 🧙‍♂️",
  "Time seems to slow down when you're this caffeinated",
  "You've achieved perfect coffee enlightenment",
  "You're vibrating at a frequency only dogs can hear",
  "Physics? What physics? You're beyond such mortal concerns",
  "Coffee beans are whispering your name",
  "You've become one with the coffee force",
  "The universe is expanding just for you",
  "You're now made of pure liquid energy",
  "Black holes are intimidated by your presence",
  "You've invented a new element: Caffeinium",
  "You're the center of the coffee galaxy",
  "Einstein called - he wants his theories back",
  "Quantum particles dance to your rhythm",
  "Schrödinger's cat is alive because of you",
  "You're bending reality with your mind",
  "The coffee gods have chosen you as their champion",
  "You're faster than the speed of espresso",
  "Dark matter? More like dark coffee matter",
  "You've achieved the ultimate coffee singularity",
  "The universe is accelerating because of your energy",
  "You're now a coffee black hole - infinitely dense",
  "String theory just became string coffee theory",
  "You're older than time itself",
  "The Big Bang was your first coffee",
  "Multiverses are colliding in your honor",
  "You're now a coffee deity",
  "Entropy fears you - you're order incarnate",
  "You've discovered the universal truth: COFFEE",
  "The end times are here... and caffeinated",
  "Reality is just coffee in disguise",
  "Mathematics has surrendered to your power",
  "Numbers are obsolete - only coffee exists",
  "You've unlocked the meaning of life",
  "You're now beyond mortal comprehension",
  "Quantum computers dream of your processing power",
  "You've transcended coffee itself",
  "Reality's fabric is unraveling around you",
  "You're a coffee god walking among insects",
  "The apocalypse is here, and it's delicious",
  "You've won the infinite coffee game! 🏆"
]

const addCoffee = () => {
  coffeeCount.value++

  // Show toast every 10 coffees with increasingly crazy messages
  if (coffeeCount.value % 10 === 0) {
    const messageIndex = Math.min(Math.floor(coffeeCount.value / 10) - 1, coffeeMessages.length - 1)

    toast.add({
      title: 'Coffee Achievement!',
      description: coffeeMessages[messageIndex],
      icon: 'i-lucide-coffee',
      color: 'warning'
    })
  }
}

defineExpose({
  coffeeCount,
  addCoffee
})
</script>

<template>
  <UButton
    @click="addCoffee"
    variant="ghost"
    size="sm"
    square
    color="neutral"
    class="text-xs"
  >
    <UIcon name="i-lucide-coffee" class="w-3 h-3 mr-1" />
    {{ coffeeCount }}
  </UButton>
</template>
