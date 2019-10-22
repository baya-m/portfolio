<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="success">
      <b-navbar-brand>FX trader</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse" class="ml-auto"></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <b-navbar-brand>
            証拠金
            <font-awesome-icon icon="database" />
          </b-navbar-brand>
          <b-nav-item href="/setting">
            Setting
            <font-awesome-icon icon="cog" />
          </b-nav-item>
          <b-nav-item href="/signout">
            Signout
            <font-awesome-icon icon="sign-out-alt" />
          </b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>