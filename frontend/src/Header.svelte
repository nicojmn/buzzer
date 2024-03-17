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
   
    let isOpen = false;
    // for testing
    let admin = { loggedIn: true };
    
    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }
</script>
{#if admin.loggedIn}
<Navbar color="light" light expand="md" container="md">
    <NavbarBrand href="/">
        <img src="img/logo_akap.png" alt="Akapella logo" width="10%">
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
{/if}