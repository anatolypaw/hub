<script lang="ts">
	import '../../app.css';
	import Header from '../../components/Header.svelte';
	import LeftMenu from '../../components/LeftMenu.svelte';
	import { page } from '$app/stores';

	let currentPath: string;
	let version: string;
	let username: string;

	// Загрузка версии из API
	const loadVersion = async () => {
		try {
			const response = await fetch('/api/about');
			if (response.ok) {
				const data = await response.json();
				version = data.version || 'unknown 1';
			} else {
				console.error('Failed to load version:', response.status, response.statusText);
				version = 'unknown 2';
			}
		} catch (error) {
			console.error('Error loading version:', error);
			version = 'unknow 3';
		}
	};

	// Загрузка имени пользователя из API
	const loadUsername = async () => {
		try {
			const response = await fetch('/api/userinfo');
			if (response.ok) {
				const data = await response.json();
				username = data.username || 'unknown 1';
			} else {
				console.error('Failed to load username:', response.status, response.statusText);
				username = 'unknown 2';
			}
		} catch (error) {
			console.error('Error loading username:', error);
			username = 'unknown 3';
		}
	};

	$: {
		$page;
		currentPath = $page.url.pathname;
	}

	loadVersion();
	loadUsername();
</script>

<div class="font-mono h-screen">
	<Header {username} />

	<div class="flex h-[calc(100vh-2.5rem)]">
		<LeftMenu {currentPath} {version} />
		<!-- Под хэдером контент -->
		<div id="content" class="bg-white p-4 overflow-auto w-full">
			<slot />
		</div>
	</div>
</div>
