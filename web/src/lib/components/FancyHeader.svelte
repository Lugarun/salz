<script lang="ts">
export let headerLevel = '1';
</script>

<div class="fancy-header">
  <div class={`icon icon-level-${headerLevel}`}>
    <slot name="icon"></slot>
  </div>
  <svelte:element this={`h${headerLevel}`}>
    <slot></slot>
  </svelte:element>
</div>

<style>
.fancy-header {
  --slant-px: 30px;
  background: lightgray;
  padding: 0.6em 2.5em 0.6em 1em;
  margin-left: -10px;
  position: relative;
  display: flex;
  overflow: visible;
  clip-path: polygon(
    0 -3em,
    calc(100% - var(--slant-px)) -3em,
    calc(100% - var(--slant-px)) 0,
    100% 0,
    calc(100% - var(--slant-px)) 100%,
    calc(100% - var(--slant-px)) 150%,
    0 150%
  );
}

.fancy-header::before {
  content: " ";
  display: block;
  margin: 0;
  position: absolute;
  top: calc(100% - 10px);
  left: -10px;
  width: 20px;
  height: 20px;
  background: gray;
  clip-path: polygon(50% 50%, 100% 100%, 100% 50%);
}

.icon {
  margin-top: -0.6em;
  margin-right: 1rem;
  line-height: 0;
}

.icon.icon-level-1 {
  font-size: 5em;
}

.icon.icon-level-2 {
  font-size: 4em;
}

.icon.icon-level-3 {
  font-size: 3em;
}

.icon.icon-level-4 {
  font-size: 2.5em;
}

.icon.icon-level-5 {
  font-size: 2em;
}

.icon.icon-level-6 {
  font-size: 1.5em;
}

div :global(h1),
div :global(h2),
div :global(h3),
div :global(h4),
div :global(h5),
div :global(h6) {
  display: inline;
  line-height: 1.5em;
}
</style>