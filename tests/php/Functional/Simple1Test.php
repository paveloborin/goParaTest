<?php
namespace functional;

class Simple1Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(1);
        $this->assertTrue(
            true
        );
    }

}