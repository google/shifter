<template>
  <div class="item pb-4">
    <div class="details">
      <h1 class="text-2xl font-bold text-shifter-red-muted">
        <slot name="title"></slot>
      </h1>
      <h3 class="text-baseline italic">
        <slot name="subtitle"></slot>
      </h3>
      <slot></slot>
    </div>
  </div>
</template>