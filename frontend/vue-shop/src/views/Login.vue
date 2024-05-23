<template>
  <div class="row">
    <div class="col s8 offset-s2" style="margin-top: 50px;">
      <form>
        <div class="input-field col s12">
          <input id="email" type="email" class="validate" required>
          <label for="email" class="active">Email</label>
        </div>
      </form>
      <form>
        <div class="input-field col s12">
          <input id="password" type="password" class="validate" required>
          <label for="password" class="active">Password</label>
        </div>
      </form>
      <div>
        <button class="btn waves-effect grey darken-3" type="submit" name="action">Login
            <i class="material-icons right"></i>
        </button>
    </div>
      <label for="email">Forgot password?</label>
    </div>
  </div>
</template>
  
<script>
</script>

<style>
    .indented{
        margin-left:100px;
    }
</style>