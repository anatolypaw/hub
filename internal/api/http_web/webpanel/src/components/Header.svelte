<script lang="ts">
	import Button from './Button.svelte';
	import dm_logo from '../static/logo.svg';

	export let username: string;
</script>

<header class="bg-white border-gray-500 border-b h-10 flex justify-between items-center">
	<!-- Логотип -->
	<div class="flex items-center pl-4">
		<img src={dm_logo} alt="Logo" class="w-6 h-6" />
		<h2 class="text-2xl font-bold text-center pl-2">МОЛОКОД</h2>
	</div>

	<!-- Текст в центре -->
	<div class="grow items-center text-center">Управление маркировкой</div>

	<!-- Кнопка выхода и отображение текущего пользователя -->
	<div class="items-center mr-4">
		<span class="text-gray-700 mr-4">{username ? username : '----'}</span>
		<Button>Выйти</Button>
	</div>
</header>
