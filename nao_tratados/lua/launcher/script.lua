function testTrigger()
    print("test fired the trigger")
end

function loopTrigger()
    print("i found a loop!")
end

print("test lua script")
waitFor("golang + lua")
trigger("test", testTrigger)
trigger("loop", loopTrigger)

