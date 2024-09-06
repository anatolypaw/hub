<script lang="ts">
	import { onMount } from 'svelte';
	import Button from '../../../components/Button.svelte';

	interface Result {
		Gtin: string;
		Desc: string;
		StoreCount: number;
		GetCodeForPrint: boolean;
		AllowProduce: boolean;
		Upload: boolean;
		Created: string;
	}

	interface Data {
		Error: string;
		Result: Result[];
	}

	let data: Data = {
		Error: '',
		Result: []
	};

	async function getGoods() {
		try {
			const response = await fetch('/api/goods');
			const result = await response.json();
			data = result;
		} catch (error) {}
	}

	onMount(getGoods);

	let newGood: Result = {
		Gtin: '',
		Desc: '',
		StoreCount: 0,
		GetCodeForPrint: false,
		AllowProduce: false,
		Upload: false,
		Created: ''
	};

	const addGood = () => {
		// Добавляем новый продукт в список
		data.Result = [...data.Result, { ...newGood, Created: new Date().toISOString() }];
		// Сбрасываем форму
		newGood = {
			Gtin: '',
			Desc: '',
			StoreCount: 0,
			GetCodeForPrint: false,
			AllowProduce: false,
			Upload: false,
			Created: new Date().toISOString()
		};
	};
</script>

<div class="">
	<table class="min-w-full divide-y divide-gray-200 text-sm">
		<thead class="bg-gray-100 sticky top-0">
			<tr>
				<th class="py-1 border">Название</th>
				<th class="py-1 border w-40">GTIN</th>
				<th class="py-1 border w-40">Запас для печати</th>
				<th class="py-1 border">Получать КМ</th>
				<th class="py-1 border">Выгружать КМ</th>
				<th class="py-1 border">Доступен</th>
				<th class="py-1 border w-40">Создан</th>
				<th class="py-1 border"> -- </th>
			</tr>
		</thead>
		<tbody class="divide-y divide-gray-200">
			{#if data.Result != null}
				{#each data.Result as item}
					<tr>
						<td class="py-1 px-2 text-left border">{item.Desc}</td>
						<td class="py-1 px-2 text-center border">{item.Gtin}</td>
						<td class="py-1 px-2 text-center border">{item.StoreCount}</td>
						<td class="py-1 px-2 text-center border">
							<input type="checkbox" checked={item.GetCodeForPrint} disabled class="h-4 w-4" />
						</td>
						<td class=" py-1 px-2 text-center border">
							<input type="checkbox" checked={item.Upload} disabled class="h-4 w-4" />
						</td>
						<td class=" py-1 px-2 text-center border">
							<input type="checkbox" checked={item.AllowProduce} disabled class="h-4 w-4" />
						</td>
						<td class=" py-1 px-2 text-right border">
							{new Date(item.Created).toLocaleString('ru-RU', {
								year: '2-digit',
								month: '2-digit',
								day: '2-digit',
								hour: '2-digit',
								minute: '2-digit'
							})}
						</td>
						<td></td>
					</tr>
				{/each}
			{/if}

			<tr class="">
				<td class="border">
					<input
						id="desc"
						type="text"
						bind:value={newGood.Desc}
						placeholder="Название"
						class="py-1 px-2 w-full"
					/>
				</td>
				<td class="border">
					<input
						id="gtin"
						type="number"
						bind:value={newGood.Gtin}
						placeholder=" GTIN"
						class="py-1 px-2 w-full text-center"
					/>
				</td>
				<td class="border">
					<input
						id="number"
						type="number"
						bind:value={newGood.StoreCount}
						placeholder="Запас для печати, шт"
						class="py-1 pl-3 w-full text-center"
					/>
				</td>
				<td class="border text-center">
					<input
						id="getForPrint"
						type="checkbox"
						bind:checked={newGood.GetCodeForPrint}
						class="py-1 px-2 w-full h-4 w-4"
					/>
				</td>
				<td class="border text-center">
					<input
						id="allowproduce"
						type="checkbox"
						bind:checked={newGood.AllowProduce}
						class="py-1 px-2 w-full h-4 w-4"
					/>
				</td>
				<td class=" border text-center">
					<input id="upload" type="checkbox" bind:checked={newGood.Upload} class="h-4 w-4" />
				</td>
				<td class=" border text-center"> </td>
				<td>
					<button on:click={addGood} class="py-1 w-full text-white bg-blue-400 rounded">
						Добавить
					</button>
				</td>
			</tr>
		</tbody>
	</table>

	{#if data.Error != ''}
		<div class="mt-4 text-red-600">
			{data.Error}
		</div>
	{/if}
</div>

<style>
	/* Можно добавить свои стили, если нужно */
</style>
