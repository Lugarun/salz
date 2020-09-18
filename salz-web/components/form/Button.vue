<template>
  <button :type="type" :title="label" :aria-label="label" @click="$emit('click')">
    <div v-if="corners" class="corner1"></div>
    <div v-if="corners" class="corner2"></div>
    <slot></slot>
  </button>
</template>

<style lang="scss" scoped>
$cornerSize: 20px;
$cornerWidth: 5px;

button {
  background: radial-gradient(var(--color-primary-1), var(--color-primary-5));
  border: 2px solid var(--color-primary);
  border-radius: 5px;
  color: var(--color-light);
  padding: 10px 15px;
  position: relative;
  cursor: pointer;

  &:hover {
    background: radial-gradient(var(--color-primary-3), var(--color-primary-7));
  }

  &.active {
    background: radial-gradient(var(--color-primary-5), var(--color-primary));
  }
}

button > div[class^='corner'] {
  position: absolute;
  background: var(--color-primary);
  width: $cornerSize;
  height: $cornerSize;
}

button > div.corner1 {
  top: 0;
  left: 0;
  clip-path: polygon(
    0 0,
    100% 0,
    ($cornerSize - $cornerWidth) $cornerWidth,
    $cornerWidth $cornerWidth,
    $cornerWidth ($cornerSize - $cornerWidth),
    0 100%
  );
}

button > div.corner2 {
  bottom: 0;
  right: 0;
  clip-path: polygon(
    0 100%,
    $cornerWidth ($cornerSize - $cornerWidth),
    ($cornerSize - $cornerWidth) ($cornerSize - $cornerWidth),
    ($cornerSize - $cornerWidth) $cornerWidth,
    100% 0,
    100% 100%
  );
}
</style>

<script>
export default {
  props: {
    type: {
      type: String,
      default: 'button',
    },
    label: {
      type: String,
      required: true,
    },
    corners: {
      type: Boolean,
      default: true,
    },
  },
};
</script>
