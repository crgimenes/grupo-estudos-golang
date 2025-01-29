function testTrigger()
    print("test fired the trigger")
end

function loopTrigger()
    print("i found a loop!")
end

function ping()
    print("pong")
end

trigger("ping", ping)

print("test lua script")

waitFor("golang + lua")
trigger("test", testTrigger)
trigger("loop", loopTrigger)

