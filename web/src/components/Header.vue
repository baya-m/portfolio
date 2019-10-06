<template>
<div>
  <b-navbar toggleable="lg" type="dark" variant="success">
    <b-navbar-brand to="#">
    Portfolio
    </b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>
      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item to="/signup">Signup</b-nav-item>   
        <b-nav-item to="/login">Login</b-nav-item>   
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</div>
</template>

<script>
export default {};
</script>

<style>
logo img{

}
</style>