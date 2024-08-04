<script lang="ts">
	export let currentPath: string = '';

	const isActive = (path: string) => currentPath === path;
</script>

<nav class="bg-white w-48 border-r border-gray-500">
	<ul>
		<li>
			<a href="/" class={`hover:bg-gray-200 block py-2 px-4 ${isActive('/') ? 'bg-gray-300' : ''}`}>
				Главная
			</a>
		</li>
		<li>
			<a href="/" class={`hover:bg-gray-200 block py-2 px-4 ${isActive('/') ? 'bg-gray-300' : ''}`}>
				О нас
			</a>
		</li>
		<li>
			<a
				href="/login"
				class={`hover:bg-gray-200 block py-2 px-4 ${isActive('/login') ? 'bg-gray-300' : ''}`}
			>
				Логин
			</a>
		</li>
	</ul>
</nav>
