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
          I'm a full-stack developer specializing in Vue.js, Node.js, and TypeScript.
          I enjoy building user-friendly applications and solving complex technical challenges.
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
        <div class="space-y-4">
          <div>
            <h3 class="font-medium">Senior Full-Stack Developer</h3>
            <p class="text-sm text-slate-600 mb-1">Tech Company • 2022 - Present</p>
            <p class="text-sm text-slate-700">
              Lead development of scalable web applications using modern JavaScript frameworks.
              Mentored junior developers and contributed to architecture decisions.
            </p>
          </div>
          <div>
            <h3 class="font-medium">Full-Stack Developer</h3>
            <p class="text-sm text-slate-600 mb-1">Startup Inc • 2020 - 2022</p>
            <p class="text-sm text-slate-700">
              Built and maintained multiple client applications using Vue.js and Node.js.
              Implemented CI/CD pipelines and improved development workflows.
            </p>
          </div>
        </div>
      </div>

      <!-- Projects -->
      <div>
        <h2 class="text-lg font-semibold mb-3">Projects</h2>
        <div class="space-y-3">
          <div>
            <h3 class="font-medium">E-commerce Platform</h3>
            <p class="text-sm text-slate-700">
              Full-stack e-commerce solution with Vue.js frontend and Node.js backend.
              Features include user authentication, payment processing, and admin dashboard.
            </p>
          </div>
          <div>
            <h3 class="font-medium">Task Management App</h3>
            <p class="text-sm text-slate-700">
              Real-time task management application with collaborative features.
              Built with React, Socket.io, and MongoDB.
            </p>
          </div>
        </div>
      </div>

      <!-- Contact -->
      <div>
        <h2 class="text-lg font-semibold mb-3">Contact</h2>
        <div class="text-sm text-slate-700">
          <p>Email: adrian@example.com</p>
          <p>LinkedIn: linkedin.com/in/adrianlatorre</p>
          <p>GitHub: github.com/adrianlatorre</p>
        </div>
      </div>
    </div>
  </div>
</template>
