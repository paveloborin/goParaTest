<?php
namespace functional;

class Simple2Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(2);
        $this->assertTrue(
            true
        );
    }

}