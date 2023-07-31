<script lang="ts">
	import { createBlocksService } from '$lib/services/blocksService';
</script>

<div class="flex flex-col w-full items-stretch">
	<div class="shadow flex items-center px-8 py-4">
		<span class="font-bold">LOGO</span>
	</div>

	<div class="px-8 py-4">		
		{#await createBlocksService().getFullList()}
			<div class="h-full w-full flex items-center justify-center">
				<span>Loading...</span>
			</div>
		{:then blocks}
			<div class="w-full flex flex-col">
				<span class="text-2xl font-medium mb-4">Blocks</span>
				<div class="columns-sm">
					{#each [...blocks, ...blocks] as block, index}
						<div class="rounded-xl shadow p-4 flex flex-col gap-3 mb-4 break-inside-avoid-column bg-white">
							<div class="w-full aspect-video bg-surface-100 rounded-xl">
	
							</div>						
							<div class="flex flex-col pl-1">
								<span class="text-xl font-bold text-surface-700">
									{block.name}
								</span>				
								<span class="text-sm text-surface-500">
									by {block.user.name}
								</span>
							</div>
							<div>
								<span class="badge variant-soft-secondary">{block.category}</span>
								<span class="badge variant-soft-tertiary">⭐ 99</span>
							</div>
							<hr>
							<div class="flex flex-row-reverse justify-between">
								<button type="button" class="btn btn-sm variant-soft-primary">
									⭐ Save in list
								</button>
							</div>
						</div>
					{/each}
				</div>
			</div>
		{/await}
	</div>
</div>
