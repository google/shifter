<template>
  <div class="item pb-4">
    <div class="details">
      <h1 class="text-lg font-bold text-shifter-red-muted">
        <slot name="title"></slot>
      </h1>
      <h3 class="text-sm italic">
        <slot name="subtitle"></slot>
      </h3>
      <slot></slot>
    </div>
  </div>
</template>