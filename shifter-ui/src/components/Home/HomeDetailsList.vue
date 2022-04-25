<script setup>
import HomeDetailsListItem from "../DetailsListItem.vue";
import DocumentationIcon from "../icons/IconDocumentation.vue";
import CommunityIcon from "../icons/IconCommunity.vue";
</script>

<template>
  <HomeDetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Documentation</template>

    Shifter's
    <a target="_blank" href="https://github.com/google/shifter">official documentation</a>
    provides you with all information you need to get started with migrating.
  </HomeDetailsListItem>

  <HomeDetailsListItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Community</template>

    Found a bug? Raise an issue on
    <a target="_blank" href="https://github.com/google/shifter/issues">GitHub</a> or alternatively
    become a contributor.
  </HomeDetailsListItem>

</template>
