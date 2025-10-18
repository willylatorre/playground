<script setup lang="ts">
// About Me page - no complex functionality needed
</script>

<template>
  <div class="max-w-4xl mx-auto py-8">
    <div class="space-y-8">
      <!-- Header -->
      <div>
        <h1 class="text-2xl font-bold mb-2">Adrian Latorre</h1>
        <p class="text-slate-600">Full-Stack Developer</p>
      </div>

      <!-- Bio -->
      <div>
        <h2 class="text-lg font-semibold mb-3">About</h2>
        <p class="text-slate-700 leading-relaxed">
          I'm a full-stack developer specializing in Vue.js, Node.js, and TypeScript. I enjoy
          building user-friendly applications and solving complex technical challenges.
        </p>
      </div>

      <!-- Skills -->
      <div>
        <h2 class="text-lg font-semibold mb-3">Skills</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div>
            <h3 class="font-medium text-sm mb-2">Frontend</h3>
            <ul class="text-sm text-slate-600 space-y-1">
              <li>Vue.js</li>
              <li>React</li>
              <li>TypeScript</li>
            </ul>
          </div>
          <div>
            <h3 class="font-medium text-sm mb-2">Backend</h3>
            <ul class="text-sm text-slate-600 space-y-1">
              <li>Node.js</li>
              <li>Go</li>
              <li>PostgreSQL</li>
            </ul>
          </div>
          <div>
            <h3 class="font-medium text-sm mb-2">Tools</h3>
            <ul class="text-sm text-slate-600 space-y-1">
              <li>Git</li>
              <li>Docker</li>
              <li>AWS</li>
            </ul>
          </div>
          <div>
            <h3 class="font-medium text-sm mb-2">Other</h3>
            <ul class="text-sm text-slate-600 space-y-1">
              <li>Problem Solving</li>
              <li>System Design</li>
              <li>Team Leadership</li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Experience -->
      <div>
        <h2 class="text-lg font-semibold mb-3">Experience</h2>
        <div class="space-y-6">
          <div>
            <h3 class="font-semibold text-lg">Adyen</h3>
            <p class="text-sm text-slate-600 mb-2">Amsterdam, August 2020 - Present</p>
            <p class="text-slate-700 mb-3">
              After a proper analysis, one of the most important companies in the Fintech sector has
              decided to work with a:
            </p>
            <ul class="text-sm text-slate-700 space-y-1 ml-4">
              <li>
                • Currently, Frontend Engineer in charge of the main Adyen.com and derivates
                websites
              </li>
              <li>• Previously, Frontend Engineer at the Billing Stream</li>
              <li>• Improve the invoicing experience for both merchants and AMs</li>
              <li>• Develop internal tools related to invoices and contracts</li>
            </ul>
            <div class="mt-4 p-4 bg-slate-50 rounded-lg">
              <h4 class="font-medium text-slate-900 mb-2">
                Special focus at improving the Adyen Demos:
              </h4>
              <ul class="text-sm text-slate-700 space-y-1 ml-4">
                <li>
                  • Create the official Adyen Demos platform to enable employees to generate demos
                  of adyen products in less than 1 minute (Nuxt 3, Aurora serverless 2)
                </li>
                <li>
                  • Integrate every Adyen product to the demo platform ( Checkout, POS, Platforms,
                  Tap-to-pay, tipping, giving, etc…)
                </li>
                <li>• Integrate Adyen MCP and other interactions with agents and genAI</li>
              </ul>
            </div>
          </div>

          <div>
            <h3 class="font-semibold text-lg">Housfy</h3>
            <p class="text-sm text-slate-600 mb-2">Barcelona, March 2019 - August 2020</p>
            <p class="text-slate-700 mb-3">First usage of a in the PropTech Sector:</p>
            <ul class="text-sm text-slate-700 space-y-1 ml-4">
              <li>• Lead the Frontend stack of Housfy (Vue.js, Nuxt.js)</li>
              <li>• Implement MyHousfy, a fully featured PWA to provide value for the customers</li>
              <li>• Implement a good testing (Jest/Cypress) environment</li>
              <li>• Implement analytics tracking</li>
            </ul>
          </div>

          <div>
            <h3 class="font-semibold text-lg">Glovoapp</h3>
            <p class="text-sm text-slate-600 mb-2">Barcelona, Sept 2015 - March 2019</p>
            <p class="text-slate-700 mb-3">
              One of the biggest companies in the delivery ecosystem right now. Glovoapp has worked
              with a as Technical Lead and has done pretty cool things with it:
            </p>
            <ul class="text-sm text-slate-700 space-y-1 ml-4">
              <li>• Lead the Frontend stack of Glovo (Vue.js, Nuxt.js, HeadlessCMS)</li>
              <li>• Lead the Growth / B2B team</li>
              <li>• Implement & A/B Test multiple RPO improvements</li>
              <li>• Implement Glovo Prime</li>
              <li>• Implement Glovo Member Get Member program</li>
              <li>• Implement Glovo Business</li>
              <li>• Migration of the Glovo webapp from Angular to Vue.js</li>
            </ul>
            <p class="text-sm text-slate-600 mt-2 italic">
              Here is a great summary of it: A wonderful journey: migrating from Angular to Vue (and
              why not React)
            </p>
          </div>

          <div>
            <h3 class="font-semibold text-lg">Oliva App</h3>
            <p class="text-sm text-slate-600 mb-2">Barcelona, Feb 2015 - Sept 2015</p>
            <p class="text-slate-700 mb-3">
              Small startup with a focus on delivering a high-tech solution to improve the fidelity
              amongst merchants. Oliva used a alone to build the whole Android application and
              website and implemented some very interesting technologies as NFC and QR codes.
            </p>
          </div>
        </div>
      </div>

      <!-- Contact -->
      <div>
        <h2 class="text-lg font-semibold mb-3">Contact</h2>
        <div class="text-sm text-slate-700">
          <p>Email: <a href="mailto:adri.latorre@gmail.com">adri.latorre@gmail.com</a></p>
          <p>
            LinkedIn:
            <a href="https://www.linkedin.com/in/adrianlatorre/" target="_blank"
              >linkedin.com/in/adrianlatorre</a
            >
          </p>
          <p>
            GitHub:
            <a href="https://github.com/willylatorre" target="_blank">github.com/willylatorre</a>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
