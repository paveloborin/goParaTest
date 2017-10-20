<?php
namespace functional;

class Simple6Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(8);
        $this->assertTrue(
            true
        );
    }

}