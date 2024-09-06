<script lang="ts">
	export let onClick: () => void;
</script>

<button
	on:click={onClick}
	class="bg-white hover:bg-gray-200 px-3 rounded border-solid border-2 border-gray-500 text-gray-700"
>
	<slot></slot>
</button>
