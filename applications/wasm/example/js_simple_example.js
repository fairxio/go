

function PerformExecutableWorkflow() {
    var x = "yes";
    console.log("FairX Version: " + fairx.version);
    fairx.ret = fairx.version;

    fairx.callParticipant("did:fairx:zkfjlsdkjfsldkfjsldkjfsd", "executable.functionName", "arg1", "arg2");
    fairx.callParticipant();

}

function PerformParticipantCalling(ident) {

    fairx.callParticipant(ident, "SomeFunction", "Nah");
    var y = "nope";

}

