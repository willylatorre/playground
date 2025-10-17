<script setup lang="ts">
import { ref } from 'vue'

const coffeeCount = ref(Math.floor(Math.random() * 15) + 5)

const toast = useToast()

// Crazy coffee messages that escalate in incredulity
const coffeeMessages = [
  "Ah, a fresh cup of coffee! ☕",
  "Second cup already? You're committed! ☕☕",
  "Three cups? You're officially a coffee enthusiast!",
  "Four cups? Your heart must be racing! 💓",
  "Five cups? Are you even human anymore?",
  "Six cups? I think your blood is now 90% caffeine!",
  "Seven cups? You're basically a coffee-powered robot now 🤖",
  "Eight cups? Time dilation is kicking in... everything feels slower",
  "Nine cups? You've achieved caffeine enlightenment 🧘",
  "Ten cups? You're now vibrating at the molecular level",
  "Eleven cups? Physics no longer applies to you ⚛️",
  "Twelve cups? You're communicating with coffee beans telepathically",
  "Thirteen cups? Unlucky? No, you're just immortal now",
  "Fourteen cups? Time is flowing backwards for you ⏰",
  "Fifteen cups? You've broken the space-time continuum",
  "Sixteen cups? Parallel universes are jealous of your energy",
  "Seventeen cups? You're now made of pure caffeine crystals 💎",
  "Eighteen cups? Black holes are afraid of your gravitational pull",
  "Nineteen cups? You've invented a new state of matter: 'Caffeinated Plasma'",
  "Twenty cups? You're now the center of the universe 🌌",
  "Twenty-one cups? Einstein's theories were wrong, you proved it",
  "Twenty-two cups? Quantum entanglement? You cause it",
  "Twenty-three cups? Schrödinger's cat is both alive AND dead because of you",
  "Twenty-four cups? You're bending spoons with your mind (and coffee)",
  "Twenty-five cups? The coffee bean gods have blessed you",
  "Twenty-six cups? You're communicating faster than light",
  "Twenty-seven cups? Dark matter is actually made of coffee grounds",
  "Twenty-eight cups? You've achieved coffee singularity",
  "Twenty-nine cups? The universe is expanding because of your energy",
  "Thirty cups? You're now a coffee black hole - nothing can escape",
  "Thirty-one cups? String theory? You just proved there are 11 dimensions",
  "Thirty-two cups? You're older than the universe itself now",
  "Thirty-three cups? Big Bang? That was just you brewing your first cup",
  "Thirty-four cups? Multiverses are colliding because of your caffeine",
  "Thirty-five cups? You're now a coffee deity - worship begins",
  "Thirty-six cups? Entropy decreases around you - you're anti-chaos",
  "Thirty-seven cups? You've discovered the theory of everything: COFFEE",
  "Thirty-eight cups? The end of the universe is nigh... or is it just more coffee?",
  "Thirty-nine cups? Reality is just a simulation running on your coffee",
  "Forty cups? Congratulations! You've broken mathematics itself",
  "Forty-one cups? Numbers no longer exist - only coffee",
  "Forty-two cups? You've answered the ultimate question: it's coffee",
  "Forty-three cups? The meaning of life is... more coffee",
  "Forty-four cups? You're now beyond comprehension",
  "Forty-five cups? Even quantum computers can't calculate your energy",
  "Forty-six cups? You've achieved coffee transcendence",
  "Forty-seven cups? The fabric of reality is unraveling",
  "Forty-eight cups? You're now a coffee god among mortals",
  "Forty-nine cups? The apocalypse is here... and it's caffeinated",
  "Fifty cups? You've won the game. Coffee Clicker complete! 🏆"
]

const addCoffee = () => {
  coffeeCount.value++

  // Show toast every 5 coffees with increasingly crazy messages
  const messageIndex = Math.min(Math.floor(coffeeCount.value / 5), coffeeMessages.length - 1)

  if (messageIndex > 0) {
    toast.add({
      title: 'Coffee Achievement!',
      description: coffeeMessages[messageIndex - 1],
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
