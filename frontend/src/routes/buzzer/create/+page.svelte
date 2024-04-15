<script lang="ts">
    import { onMount } from "svelte";

    let teamName = "";
    let isNameTaken = false;
    let isTooLong = false;
    let teamsName:any = [];

    onMount(async () => {
        const response = await fetch("/admin/teams");
        const teams:any = await response.json();
        teamsName = teams.map((team:any) => team.Name.toLowerCase().trim());
    })

    $: console.log(`teamsName: ${teamsName}`);
    $: console.log(`isNameTaken: ${isNameTaken}`);

    $: isNameTaken = teamsName.includes(teamName.toLowerCase().trim());
    $: teamName.length > 24 ? isTooLong = true : isTooLong = false;


</script>


{#if isNameTaken}
    <div role="alert" class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span class="text-2xl">Erreur : le nom de l'équipe est déjà utilisé !</span>
    </div>
{/if}
{#if isTooLong}
    <div role="alert" class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span class="text-2xl">Erreur : le nom de l'équipe est trop long !</span>
    </div>
{/if}

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content flex-col lg:flex-row-reverse">
    <div class="text-center lg:text-left">
      <h1 class="text-5xl font-bold">Création d'équipe</h1>
    </div>
    <div class="card shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
      <form class="card-body" method="post">
        <div class="form-control">
          <label class="label" for="teamName">
            <span class="label-text">Nom de l'équipe</span>
          </label>
          <input type="text" id="teamName" name="teamName" placeholder="Un nom original" class="input input-bordered" maxlength="24" required />
        </div>
        <div class="form-control mt-6">
          <button class="btn btn-primary" formmethod="post">Créer l'équipe</button>
        </div>
      </form>
    </div>
  </div>
</div>