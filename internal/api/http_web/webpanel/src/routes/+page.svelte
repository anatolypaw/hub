<script lang="ts">
	import { page } from '$app/stores';
	import Header from '../components/Header.svelte';
	import LeftMenu from '../components/LeftMenu.svelte';

	let currentPath: string;

	$: {
		$page;
		currentPath = $page.url.pathname;
	}
</script>

<div class="font-mono h-screen">
	<Header username="Администратор" />

	<div class="flex h-[calc(100vh-2.5rem)]">
		<LeftMenu {currentPath} />
		<!-- Под хэдером контент -->
		<div id="content" class="bg-white p-4 overflow-auto w-full">
			<h1>test</h1>
		</div>
	</div>
</div>
