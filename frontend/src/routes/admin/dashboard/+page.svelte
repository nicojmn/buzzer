<script lang="ts">
    import GameState from "$lib/components/GameState.svelte"
    import { onMount } from "svelte";

    let teams:any = []

    onMount(async () => {
        const response = await fetch("/admin/teams")
        teams = await response.json()
    })
</script>
<div> 
    <GameState />

    <div class="justify-center flex">
        <div class=" card-full mt-2">
            <button class="btn-blue">
                Bonne réponse
            </button>
            <button class="btn-blue">
                Mauvaise réponse
            </button>
        </div>
    </div>

    <div class="flex w-full">
        <table>
            <thead>
                <tr>
                    <th>Team</th>
                    <th>Score</th>
                </tr>
            </thead>
            <tbody>
                {#each teams as team}
                    <tr>
                        <th>{team.Name}</th>
                        <th>{team.Score}</th>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>
