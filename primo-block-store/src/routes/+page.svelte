<script lang="ts">
	import { createBlocksService } from '$lib/services/blocksService';
</script>

<div class="flex flex-col w-full items-stretch">
	<div class="shadow flex items-center p-4">
		<span class="font-bold">LOGO</span>
	</div>

	{#await createBlocksService().getFullList()}
		<div class="h-full w-full flex items-center justify-center">
			<span>Loading...</span>
		</div>
	{:then blocks}
		<div class="flex flex-wrap gap-4">
			{#each blocks as block}
				<div>
					1					
				</div>
			{/each}
		</div>
	{/await}
</div>
