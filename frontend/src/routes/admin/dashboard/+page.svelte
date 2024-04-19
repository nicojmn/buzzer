<script lang="ts">
    import { formatUnixTimestamp } from "$lib/utils/convertTimestamp.js"
    import { onMount } from "svelte";

    let teams:any = []
    let sortedTeams:any = []
    let podium:any = []
    let i:number = 0

    let isLocked:boolean
    let buzzer_number:number = 0

    onMount(async () => {
        const response = await fetch("/admin/teams")
        teams = await response.json()
        connect()
        lockState()
        buzzer_number = await numberLocked()
    })
    $: if (teams.length > 0) {
        sortedTeams = teams.sort((a: any, b: any) => b.PressedAt - a.PressedAt);
        podium = teams.sort((a: any, b: any) => b.Score - a.Score).slice(0, 3);
    }

    $: buzzer_number

    function connect() {
        const ws = new WebSocket("ws://localhost:8080/admin/ws")
        ws.onopen = () => {
            console.log(" ws : Connected")
            ws.send(JSON.stringify({ type: 'discover', data: 'Hello server! Dashboard page here!'}));
        }
        ws.onmessage = (event) => {
            const message = JSON.parse(event.data);

            if (message.type === 'scoreUpdate') {
                console.log('Score update received:', message.data);
                updateScoreUI(message.data.team_id, message.data.score);
                ws.send(JSON.stringify({ type: 'ack', data: 'Score update received!'}));
            }
        }
        ws.onclose = () => {
            console.log("ws : Disconnected")
        }

    }

    function nextTeam() {
        i++
    }

    function incrementScore(teamID: number) {
        try {
            fetch(`/admin/teams/${teamID}/increment`, {
            method: "POST"
        })
        } catch (error) {
            console.error(error)
        }
    }

    function updateScoreUI(teamID: number, score: number) {
        // BUG : find why it increments the score by 2 sometimes
        teams = teams.map((team: { ID: number; }) => {
            if (team.ID === teamID) {
                return { ...team, Score: score };
            }
            return team;
        });
        
    }

    async function lockState() {
        try {
            const response = await fetch("/admin/lock-state")
            if (!response.ok) {
                alert("Erreur lors de la récupération de l'état des buzzers, rechargez la page (fetch error)")
                return
            }
            let data = await response.json()
            isLocked = data.state
        } catch (error) {
            alert("Erreur lors de la récupération de l'état des buzzer, réessayez")
        }
    }       

    function lockBuzzer() {
        try {
            fetch("/admin/lock", {
                method: "POST"
            })
        } catch (error) {
            alert("Erreur lors du verrouillage du buzzer, réessayez")
        }
    }

    function unlockBuzzer() {
        try {
            fetch("/admin/unlock", {
                method: "POST"
            })
        } catch (error) {
            alert("Erreur lors du déverrouillage du buzzer, réessayez")
        }
    }

    async function numberLocked(){
        const response = await fetch("/admin/buzzer-number")
        if (!response.ok) {
            alert("Erreur lors de la récupération du nombre de buzzers pressés, rechargez la page (fetch error)")
            return
        }
        let data = await response.json()
        return data.number;
    }
        
</script>
<div> 
    
    <!-- Game state -->
    <div class="flex flex-row justify-around my-4">
        <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
            <div class="card-body">
                {#if sortedTeams.length > 0}
                    <span class="text-center"> Buzzers pressés : {buzzer_number}/{sortedTeams.length}</span>
                {:else}
                    <span class="text-center"> Buzzers pressés : loading...</span>
                {/if}
            </div>
        </div>
        

        <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
            <div class="card-body">
                <span class="text-center">État des buzzers : {isLocked ? "verrouillés" : "déverrouillés"}</span><br>
                {#if isLocked}
                    <button class="btn btn-outline btn-success mt-1" on:click={() => unlockBuzzer()}>Déverrouiller</button>
                {:else}
                    <button class="btn btn-outline btn-error mt-1" on:click={() => lockBuzzer()}>Verrouiller</button>
                {/if}
            </div>
        </div>
    
        <div class="card w-1/4 bg-white border-2 border-gray-300 rounded-lg shadow">
            <div class="card-body text-center">
                {#if sortedTeams.length > 0}
                    {#each {length : 3} as _, i}
                        <span>Team {i + 1} : {podium[i].Name}</span><br>
                    {/each}
                {:else}
                    <span>Podium loading...</span>
                {/if}
            </div>
        </div>
    </div>

    <!-- Score card and scoreboard -->
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
                    <button class="btn btn-outline btn-success" on:click={() => incrementScore(sortedTeams[i % sortedTeams.length].ID)}>
                        Bonne réponse
                    </button>
                    <button class="btn btn-outline btn-error" on:click={nextTeam}>
                        Mauvaise réponse
                    </button>
                </div>
            </div>
        </div>
    </div>

    <div class="overflow-x-auto mt-2 justify-center">
        <table class="table">
            <thead>
                <tr>
                    <th>Team</th>
                    <th>Score</th>
                    <th>Buzzer pressé </th>
                    <th><button class="btn btn-outline btn-info">Trier</button></th>
                </tr>
            </thead>
            <tbody>
                {#each teams as team}
                    <tr class="hover">
                        <th>{team.Name}</th>
                        <th>{team.Score}</th>
                        <th>{formatUnixTimestamp(team.PressedAt)}</th>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>
