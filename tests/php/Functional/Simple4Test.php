<?php
namespace functional;

class Simple4Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(5);
        $this->assertTrue(
            true
        );
    }

}