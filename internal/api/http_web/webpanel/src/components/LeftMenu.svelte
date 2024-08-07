<script lang="ts">
	export let currentPath: string = '';
	export let version: string;

	const isActive = (path: string) => currentPath === path;
</script>

<nav class="bg-white w-48 border-r border-gray-500 flex flex-col">
	<!-- Меню -->
	<ul>
		<li>
			<a href="/" class={`hover:bg-gray-200 block py-2 px-4 ${isActive('/') ? 'bg-gray-300' : ''}`}>
				Главная
			</a>
		</li>
		<li>
			<a
				href="/login.html"
				class={`hover:bg-gray-200 block py-2 px-4 ${isActive('/login') ? 'bg-gray-300' : ''}`}
			>
				Логин
			</a>
		</li>
	</ul>

	<!-- Сервисная информация -->
	<div class="mt-auto border-t border-gray-300 flex justify-center">
		<a href="/changelog" class="text-center underline">{version} </a>
	</div>
</nav>
