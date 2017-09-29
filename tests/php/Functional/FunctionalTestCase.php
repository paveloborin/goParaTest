<?php

namespace functional;

use PHPUnit\Framework\TestCase;

class FunctionalTestCase extends TestCase
{
    use \PHPUnit_Extensions_Database_TestCase_Trait;

    /**
     * @var \PDO
     */
    static private $pdo = null;
    static private $isDataSetLoaded = null;
    static private $isDataSetStaticLoaded = null;
    private $conn = null;
    private $pdoData = [];

    public function createPDODataFromGlobals()
    {
        $connection = [
            0 => ['dsn' => $GLOBALS['DB_DSN'] . ':' . $GLOBALS['DB_PORT'], 'user' => $GLOBALS['DB_USER'], 'pass' => $GLOBALS['DB_PASSWD']],
            1 => ['dsn' => "mysql:host=localhost:3306", 'user' => 'root', 'pass' => 1],
            2 => ['dsn' => "mysql:host=localhost:3307", 'user' => 'root', 'pass' => 2],
            3 => ['dsn' => "mysql:host=localhost:3308", 'user' => 'root', 'pass' => 3]
        ];
        if (isset($GLOBALS['argv'][4])) {
            $i = (int)trim(str_replace('tokenNum=', '', $GLOBALS['argv'][4]));
            $this->pdoData = $connection[$i];
        } else {
            $this->pdoData = ['dsn' => $GLOBALS['DB_DSN'] . ':' . $GLOBALS['DB_PORT'], 'user' => $GLOBALS['DB_USER'], 'pass' => $GLOBALS['DB_PASSWD']];

        }
    }

    final public function getConnection()
    {
        $this->createPDODataFromGlobals();

        if ($this->conn === null) {
            if (self::$pdo == null) {

                self::$pdo = new \PDO($this->pdoData["dsn"], $this->pdoData["user"], $this->pdoData["pass"]);

                if ($GLOBALS['DB_SCHEMA'] != 1) {
                    $this->execSchema();
                }

            }
            $this->conn = $this->createDefaultDBConnection(self::$pdo, $GLOBALS['DB_DBNAME']);
        }

        return $this->conn;

    }

    public function getDataSet()
    {
        $ds = new \PHPUnit_Extensions_Database_DataSet_CompositeDataSet([]);
        if (!self::$isDataSetStaticLoaded) {
            self::$isDataSetStaticLoaded = true;
            $ds = $this->createCompositeDataSetFromDirectory($this->getFixtureStaticPath(), $ds);
        }
        if (!self::$isDataSetLoaded) {
            self::$isDataSetLoaded = true;
            $ds = $this->createCompositeDataSetFromDirectory($this->getFixturePath(), $ds);
        }
        return $ds;
    }

    private function execSchema()
    {
        self::$pdo->exec($this->getSchemaQuery());
    }


    private function createCompositeDataSetFromDirectory($directory, \PHPUnit_Extensions_Database_DataSet_CompositeDataSet $ds)
    {
        foreach (glob($directory . '*.xml') as $file) {
            $ds->addDataSet($this->createMySQLXMLDataSet($file));
        }

        return $ds;
    }


    private function getFixturePath()
    {
        return __DIR__ . '/fixtures/';
    }

    private function getFixtureStaticPath()
    {
        return __DIR__ . '/fixturesStatic/';
    }

    private function getSchemaQuery()
    {
        return file_get_contents(__DIR__ . '/schema.sql');
    }
}