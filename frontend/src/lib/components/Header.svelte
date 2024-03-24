<script lang="ts">
    import {
    Styles,
    Collapse,
    NavbarToggler,
    NavbarBrand,
    Nav,
    Navbar,
    NavItem,
    NavLink,
    Dropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem
  } from '@sveltestrap/sveltestrap';

  import logo_akap from '$lib/images/logo_akap.png'
   
    let isOpen = false;
    
    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }
</script>
<svelte:head>
  <meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>N'oubliez pas les paroles !</title>
  <link rel="icon" type="image/png" href="favicon.png">
  <!-- Bootstrap v5.3.3 -->
  <link rel="stylesheet" href="bootstrap.min.css">
</svelte:head>

<Navbar color="light" light expand="md" container="md">
    <NavbarBrand href="/">
        <img src={logo_akap} alt="Akapella logo" width="10%">
    </NavbarBrand>
    <NavbarToggler on:click={() => (isOpen = !isOpen)} />
    <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
      <Nav class="ms-auto" navbar>
        <NavItem>
          <NavLink href="#">Dashboard</NavLink>
        </NavItem>
        <NavItem>
          <NavLink href="#">Buzzers</NavLink>
        </NavItem>
        <Dropdown nav inNavbar>
          <DropdownToggle nav caret>Options</DropdownToggle>
          <DropdownMenu end>
            <DropdownItem>Option 1</DropdownItem>
            <DropdownItem>Option 2</DropdownItem>
            <DropdownItem divider />
            <DropdownItem>Reset</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </Nav>
    </Collapse>
</Navbar>