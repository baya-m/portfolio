<template>
  <div class="light-blue-skin">
    <mdb-side-nav slim :isCollapsed="collapsed" name="Anna Deynah" :OpenedFromOutside="true" logo="https://mdbootstrap.com/img/Photos/Avatars/img%20(10).jpg" logoRound sideNavClass="sn-bg-1" mask="strong" hidden>
        <mdb-side-nav-nav>
          <mdb-side-nav-cat name="Submit blog" icon="chevron-right">
            <mdb-side-nav-item href="#" slimIcon="address-book">Submit listing</mdb-side-nav-item>
            <mdb-side-nav-item href="#" slimIcon="address-card">Registration form</mdb-side-nav-item>
          </mdb-side-nav-cat>
          <mdb-side-nav-cat name="Instruction" far icon="hand-pointer">
            <mdb-side-nav-item href="#" slimIcon="blog">For bloggers</mdb-side-nav-item>
            <mdb-side-nav-item href="#" slimIcon="pen">For authors</mdb-side-nav-item>
          </mdb-side-nav-cat>
          <mdb-side-nav-cat name="About" icon="eye">
            <mdb-side-nav-item href="#" slimIcon="handshake">Introduction</mdb-side-nav-item>
            <mdb-side-nav-item href="#" slimIcon="tasks">Monthly meetings</mdb-side-nav-item>
          </mdb-side-nav-cat>
          <mdb-side-nav-cat name="Contact me" far icon="envelope">
            <mdb-side-nav-item href="#">FAQ</mdb-side-nav-item>
            <mdb-side-nav-item href="#" slimIcon="paper-plane" >Write a message</mdb-side-nav-item>
          </mdb-side-nav-cat>
          <mdb-side-nav-item icon="angle-double-left" open-icon="angle-double-right" @toggle="toggleSlim" toggler :collapsed="collapsed" fixed>Minimize</mdb-side-nav-item>
        </mdb-side-nav-nav>
    </mdb-side-nav>
  </div>
</template>

<script>
export default {

}
</script>

<style>

</style>