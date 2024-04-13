<script lang="ts">
    import GameState from "$lib/components/GameState.svelte"
    import { onMount } from "svelte";

    let teams:any = []
    let sortedTeams:any = []
    let i:number = 0

    onMount(async () => {
        const response = await fetch("/admin/teams")
        teams = await response.json()
    })
    $: if (teams.length > 0) {
        sortedTeams = teams.sort((a: any, b: any) => b.PressedAt - a.PressedAt);
    }

    function nextTeam() {
        i++
    }


</script>
<div> 
    <GameState teamsNumber={Number(teams.length)}/>

    <div class="justify-center flex">
        <div class="card w-11/12 bg-white border-2 border-gray-300 rounded-lg shadow">
            <div class="card-body">
                {#if sortedTeams.length > 0}
                    <h2 class="card-title">Team : {sortedTeams[i % sortedTeams.length].Name}</h2>
                {:else}
                    <h2 class="card-title">Team : loading...</h2>
                {/if}
                <hr class="border rounded mt-2 mb-4">
                <div class="flex justify-evenly">
                    <button class="btn btn-outline btn-success">
                        Bonne réponse
                    </button>
                    <button class="btn btn-outline btn-error" on:click={nextTeam}>
                        Mauvaise réponse
                    </button>
                </div>
            </div>
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
