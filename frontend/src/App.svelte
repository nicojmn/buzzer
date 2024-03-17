<script lang="ts">
    import { Button } from '@sveltestrap/sveltestrap';
  import Footer from './Footer.svelte';
  import Header from './Header.svelte';
</script>

<svelte:head>
  <title>N'oubliez pas les paroles !</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="icon" type="image/png" href="img/logo_akap.png">
  <!-- Bootstrap v5.3.3 -->
  <link rel="stylesheet" href="bootstrap.min.css">
</svelte:head>

<main>
  <Header />

  <center>
    <Button color="primary" size="lg" value="" >
      Ceci est un buzzer
    </Button>
  </center>
  <Footer />
</main>

<style>
</style>
