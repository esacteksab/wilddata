<script>
    // create a prop
    export let org;
</script>


<foo class="py-4 text-lg">

        Name: <b><a href="orgs/{org.Name}">{org.Name}</a></b> <br/>
        Org ID: <b>{org.ID}</b> <br/>
        EMail: <b>{org.EMail}</b> <br/>
<br/>
    <hr />
</foo>