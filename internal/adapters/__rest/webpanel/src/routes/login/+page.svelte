<script lang="ts">
	let username: string = '';
	let password: string = '';
	let error: string = '';

	const handleLogin = async (event: Event) => {
		event.preventDefault();

		try {
			// Отправка POST-запроса на сервер
			const response = await fetch('/api/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password })
			});

			// if (!response.ok) {
			// 	// Обработка ошибок, если запрос не удался
			// 	throw new Error(response.status + ' ' + response.statusText);
			// }

			const data = await response.json();

			if (data.authToken) {
				// Сохранение authToken в куки
				document.cookie = `authToken=${data.authToken}; path=/;`;

				// Перенаправление на главную страницу
				window.location.href = '/';
			} else {
				error = data.message;
			}
		} catch (err: unknown) {
			if (err instanceof Error) {
				error = err.message;
			} else {
				error = 'Неизвестная ошибка';
			}
		}
	};
</script>

<main class="flex items-center justify-center h-screen font-mono">
	<div class="bg-white border-2 border-gray-300 rounded-lg p-4 w-full max-w-sm">
		<h2 class="text-2xl font-bold mb-2 text-center">МОЛОКОД</h2>
		<h2 class="text-1xl mb-2 text-center">Авторизация</h2>

		<form on:submit={handleLogin}>
			<div class="mb-4">
				<label class="block text-sm font-bold mb-2" for="username">Логин</label>
				<input
					id="username"
					bind:value={username}
					type="text"
					class="appearance-none border-2 border-gray-300 rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
					placeholder="Введите логин"
				/>
			</div>

			<div class="mb-6">
				<label class="block text-sm font-bold mb-2" for="password">Пароль</label>
				<input
					id="password"
					bind:value={password}
					class="appearance-none border-2 border-gray-300 rounded w-full py-2 px-3 mb-2 leading-tight focus:outline-none focus:shadow-outline"
					type="password"
					placeholder="Введите пароль"
				/>
			</div>

			<button
				class="bg-white hover:bg-gray-200 px-3 rounded border-solid border-2 border-gray-500"
				type="submit"
				>Войти
			</button>
			{#if error}
				<span class="p-2 text-red-500">{error}</span>
			{/if}
		</form>
	</div>
</main>
